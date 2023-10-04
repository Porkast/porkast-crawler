FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV env                     prod
ENV WORKDIR                 /app
ADD resource                $WORKDIR/
ADD ./bin/linux_amd64/porkast-crawler $WORKDIR/porkast-crawler
RUN chmod +x $WORKDIR/porkast-crawler

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./porkast-crawler
