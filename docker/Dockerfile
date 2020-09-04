FROM ubuntu
WORKDIR /
ADD go_build_go_ci__linux /
RUN chmod +x /go_build_go_ci__linux

ENTRYPOINT ["/go_build_go_ci__linux"]