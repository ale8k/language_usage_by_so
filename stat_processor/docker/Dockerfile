FROM node:14

WORKDIR /usr/src/app

COPY ../ .

RUN npm i
RUN npm i typescript -g
RUN npm run build


EXPOSE 3000

CMD ["node", "dist/Main.js"]