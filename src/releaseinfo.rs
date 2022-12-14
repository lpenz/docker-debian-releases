// Copyright (C) 2021 Leandro Lisboa Penz <lpenz@lpenz.org>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

use chrono::{DateTime, FixedOffset};
use color_eyre::eyre::eyre;
use color_eyre::Result;
use std::collections::HashMap;
use tracing::instrument;
use url::Url;

#[derive(Debug, Default)]
pub struct ReleaseInfo {
    pub origin: String,
    pub label: String,
    pub suite: String,
    pub version: String,
    pub date: DateTime<FixedOffset>,
    pub codename: String,
    // pub url: Vec<String>,           // []string
    pub architectures: Vec<String>, // []string
}

pub mod parser {
    use nom::bytes::complete as bytes;
    use nom::character::complete as character;
    // use nom::combinator;
    use nom::multi;
    use nom::IResult;
    use std::io::BufRead;

    use super::*;

    fn kv(input: &str) -> IResult<&str, (String, String)> {
        let (input, key) = character::alpha1(input)?;
        let (input, _) = bytes::tag(": ")(input)?;
        let (input, val) = character::not_line_ending(input)?;
        let (input, _) = character::newline(input)?;
        Ok((input, (String::from(key), String::from(val))))
    }

    fn all(input: &str) -> IResult<&str, Vec<(String, String)>> {
        let (input, kvs) = multi::many1(kv)(input)?;
        let (input, _) = bytes::tag("MD5Sum:\n")(input)?;
        Ok((input, kvs))
    }

    pub fn parse(mut bufin: impl BufRead) -> Result<HashMap<String, String>> {
        let mut input = String::default();
        bufin.read_to_string(&mut input)?;
        let result = all(&input)
            .map_err(|e| eyre!("error reading input: {:?}", e))?
            .1;
        Ok(result.into_iter().collect())
    }
}

#[instrument(fields(%url))]
pub async fn get_releaseinfo(url: &Url) -> Result<ReleaseInfo> {
    let release_url = url.join("Release")?;
    let result = reqwest::get(release_url.as_str()).await?.text().await?;
    let map = parser::parse(result.as_bytes())?;
    eprintln!("{:?}", map);
    Ok(ReleaseInfo {
        origin: map
            .get("Origin")
            .cloned()
            .ok_or_else(|| eyre!("Origin not found"))?,
        label: map
            .get("Label")
            .cloned()
            .ok_or_else(|| eyre!("Label not found"))?,
        suite: map
            .get("Suite")
            .cloned()
            .ok_or_else(|| eyre!("Suite not found"))?,
        version: map
            .get("Version")
            .cloned()
            .ok_or_else(|| eyre!("Version not found"))?,
        date: map
            .get("Date")
            .and_then(|date| DateTime::parse_from_rfc2822(&date.replace("UTC", "GMT")).ok())
            .ok_or_else(|| eyre!("Date not found"))?,
        codename: map
            .get("Codename")
            .cloned()
            .ok_or_else(|| eyre!("Codename not found"))?,
        // url: Vec<String>,           // []string
        architectures: map
            .get("Architectures")
            .ok_or_else(|| eyre!("Architectures not found"))?
            .split(' ')
            .map(String::from)
            .collect::<Vec<_>>(),
    })
}

#[test]
fn test_parse() -> Result<()> {
    const EXAMPLE: &str = "Origin: Debian
Label: Debian
Suite: oldoldstable-proposed-updates
Version: 6.0-updates
Codename: squeeze-proposed-updates
Date: Thu, 24 Sep 2015 06:14:12 UTC
Architectures: amd64 armel i386 ia64 kfreebsd-amd64 kfreebsd-i386 mips mipsel powerpc s390 sparc
Components: main contrib non-free
Description: Proposed Updates for Debian 6.0 - Not Released
MD5Sum:
 a4b4567b7790d8e87ae021d70b7c4589  1170698 Contents-amd64
 9e78cc119738c99da37bba33e6bc8506    63186 Contents-amd64.gz
 069958538fe46d0458aa64acd6e15bb0  1169747 Contents-armel
 5810269b85c03e32b15527a99b5196de    63136 Contents-armel.gz
 1f47f14d15340df7159ea4c5ef419606  1170982 Contents-i386
 d6ac0555903fbcfd1cf4da075a02ef13    63228 Contents-i386.gz
 2cd8618bb3cc6a482162c662dbcc62ab  1169047 Contents-ia64
 23a3db386037c1c67d30f6ef39b1cb09    63065 Contents-ia64.gz
 43ab20c2ebf40a83231a3f7b26d45e7b  1101735 Contents-kfreebsd-amd64
 bfb274ea0a1d91438db1d46352fd9973    57977 Contents-kfreebsd-amd64.gz
";
    assert!(matches!(parser::parse(EXAMPLE.as_bytes()), Ok(_)));
    Ok(())
}
