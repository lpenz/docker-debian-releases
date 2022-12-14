// Copyright (C) 2021 Leandro Lisboa Penz <lpenz@lpenz.org>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

use clap::Parser;
use color_eyre::eyre::eyre;
use color_eyre::Result;
use select::document::Document;
use select::predicate::Name;
use tracing::instrument;
use url::Url;

use docker_debian_releases::releaseinfo;

#[instrument]
async fn get_links(url: &Url) -> Result<Vec<Url>> {
    eprintln!("url {:?}", url.as_str());
    let res = reqwest::get(url.as_str()).await?.text().await?;
    Document::from(res.as_str())
        .find(Name("a"))
        .filter_map(|n| n.attr("href"))
        .map(|link| url.join(link).map_err(|e| eyre!(e)))
        .collect()
}

#[instrument]
async fn get_aptmirrors_releaseinfos() -> Result<()> {
    let mirrors: [Url; 1] = [
        Url::parse("http://archive.debian.org/debian/")?,
        // Url::from("http://deb.debian.org/debian"),
        // Url::from("http://ports.ubuntu.com/ubuntu-ports"),
        // Url::from("http://archive.ubuntu.com/ubuntu"),
        // Url::from("http://old-releases.ubuntu.com/ubuntu"),
        // Url::from("http://archive.raspbian.org/raspbian"),
        // Url::from("http://raspbian.raspberrypi.org/raspbian"),
        // Url::from("http://archive.raspberrypi.org/debian"),
        // Url::from("http://deb.devuan.org/merged"),
        // Url::from("http://deb.devuan.org/devuan"),
    ];
    for mirror in mirrors {
        let links = get_links(&mirror.join("dists/")?).await?;
        for link in links {
            eprintln!("link {}", link);
            if let Ok(ri) = releaseinfo::get_releaseinfo(&link).await {
                eprintln!("{:?}", ri);
            }
        }
    }
    Ok(())
}

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
pub struct Args {
    #[arg(short = 'o')]
    output: Option<std::path::PathBuf>,
}

#[tokio::main]
async fn main() -> Result<()> {
    color_eyre::install()?;
    env_logger::init();
    let _args = Args::parse();
    get_aptmirrors_releaseinfos().await.unwrap();
    Ok(())
}
