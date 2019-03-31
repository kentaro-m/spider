FROM node:10

WORKDIR /src
COPY . /src
RUN npm install --production
EXPOSE 3000
CMD ["npm", "start"]