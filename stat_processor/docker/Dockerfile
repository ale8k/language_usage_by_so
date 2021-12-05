FROM node:14

WORKDIR /usr/src/app

COPY . .

RUN npm i
RUN npm run build


EXPOSE 3000

CMD ["node", "dist/Server.js"]