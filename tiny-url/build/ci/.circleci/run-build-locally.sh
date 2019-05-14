#!/usr/bin/env bash
curl --user be4408e7b9002d3bad1c654f5156154efd71c6ae: \
    --request POST \
    --form revision=6dcb6c46edcb8609eefd278eb4765025a350f612 \
    --form config=@config.yml \
    --form notify=false \
        https://circleci.com/api/v1.1/project/edenriquez/system-designs
