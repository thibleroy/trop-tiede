FROM node:alpine
RUN mkdir /service
COPY . /service
WORKDIR /service
RUN npm install
CMD ["npm", "run", "dev"]
