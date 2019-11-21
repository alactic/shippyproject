FROM alpine
ADD shippyproject-srv /shippyproject-srv
ENTRYPOINT [ "/shippyproject-srv" ]
