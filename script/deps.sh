#!/usr/bin/env bash

a="""

github.com/go-kit/kit/log
github.com/go-logfmt/logfmt

github.com/go-stack/stack
github.com/pkg/errors


github.com/ethereum/go-ethereum
github.com/json-iterator/go
github.com/modern-go/reflect2
github.com/modern-go/concurrent


github.com/satori/go.uuid
github.com/yanyiwu/gojieba

github.com/tendermint/tmlibs


github.com/primasio/go-base36

"""

for i in $a; do
    if [[ ${i:0:1} != "#" ]];then
    gopm get -l $i
    fi
done