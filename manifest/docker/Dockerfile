FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV env                     prod
ENV WORKDIR                 /app
ADD resource                $WORKDIR/
ADD ./bin/linux_amd64/guoshaofm-crawler $WORKDIR/guoshaofm-crawler
RUN chmod +x $WORKDIR/guoshaofm-crawler

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./guoshaofm-crawler
