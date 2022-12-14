# Copyright (C) 2019 Leandro Lisboa Penz <lpenz@lpenz.org>
# This file is subject to the terms and conditions defined in
# file 'LICENSE', which is part of this source code package.

import os

pjoin = os.path.join

if False:
    Environment = None

env = Environment(ENV=os.environ)

env.Command("_apt-mirrors.json", [], "cargo run --bin apt-mirrors-info -- -o $TARGET")

# env.Command(
#     "_github-branches.json", ["github-branch-workflow-run"], "./$SOURCE $TARGET"
# )


# def jsonRender(target, source, jsons):
#     env.Command(
#         target,
#         source,
#         " ".join(
#             ["./json-tmpl-render"]
#             + ["--json " + j for j in jsons]
#             + ["$SOURCE", "$TARGET"]
#         ),
#     )
#     env.Depends(target, jsons)
#     env.Depends(target, ["json-tmpl-render"])


# jsonRender(
#     "index.md", "index.md.tmpl", jsons=["_apt-mirrors.json", "_github-branches.json"]
# )
# jsonRender(
#     "git-update-image-branches",
#     "git-update-image-branches.tmpl",
#     jsons=["_apt-mirrors.json"],
# )

# for rustbin in [
#     "apt-mirrors-info",
#     "bring-random-branch",
#     "dockerhub-set-descriptions",
#     "docker-manifest-set-arch",
#     "github-branch-workflow-run",
#     "json-tmpl-render",
#     "spurious-branches",
# ]:
#     env.Command(
#         gobase,
#         [pjoin("go/cmd", gobase, "main.go"), "go/common/common.go"],
#         "go build " + pjoin("./go/cmd", gobase),
#     )

# for gobase in [
#     "apt-mirrors-info",
#     "bring-random-branch",
#     "dockerhub-set-descriptions",
#     "docker-manifest-set-arch",
#     "github-branch-workflow-run",
#     "json-tmpl-render",
#     "spurious-branches",
# ]:
#     env.Command(
#         gobase,
#         [pjoin("go/cmd", gobase, "main.go"), "go/common/common.go"],
#         "go build " + pjoin("./go/cmd", gobase),
#     )
