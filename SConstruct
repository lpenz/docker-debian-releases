# Copyright (C) 2019 Leandro Lisboa Penz <lpenz@lpenz.org>
# This file is subject to the terms and conditions defined in
# file 'LICENSE', which is part of this source code package.

import os

pjoin = os.path.join

if False:
    Environment = None

env = Environment(ENV=os.environ)

env.Command('_apt-mirrors.json', ['apt-mirrors-info'], './$SOURCE $TARGET')
env.Command('_travis-branch-jobs.json',
            ['_apt-mirrors.json', 'travis-branch-jobs'],
            './travis-branch-jobs $SOURCE $TARGET')


def jsonRender(target, source, jsons):
    env.Command(
        target, source,
        ' '.join(['./json-tmpl-render'] + ['--json ' + j for j in jsons] +
                 ['$SOURCE', '$TARGET']))
    env.Depends(target, jsons)
    env.Depends(target, ['json-tmpl-render'])


jsonRender(
    'index.md',
    'index.md.tmpl',
    jsons=['_apt-mirrors.json', '_travis-branch-jobs.json'])
jsonRender(
    'git-update-image-branches',
    'git-update-image-branches.tmpl',
    jsons=['_apt-mirrors.json'])

for gobase in [
        'apt-mirrors-info',
        'bring-random-branch',
        'json-tmpl-render',
        'travis-branch-jobs',
        'dockerhub-set-descriptions',
]:
    env.Command(gobase,
                [pjoin('go/cmd', gobase, 'main.go'), 'go/common/common.go'],
                'go build ' + pjoin('./go/cmd', gobase))
