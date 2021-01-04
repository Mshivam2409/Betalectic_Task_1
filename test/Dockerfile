FROM node:latest
RUN mkdir /test
WORKDIR /test
COPY ./* ./
EXPOSE 5000
RUN yarn install
RUN yarn run tsc 
CMD [ "node","index.js" ]
