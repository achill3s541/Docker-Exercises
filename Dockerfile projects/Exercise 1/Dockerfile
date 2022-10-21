FROM node:19.0.0
WORKDIR /usr/app
COPY package.json package-lock.json ./
RUN npm install
COPY src ./src
EXPOSE 3030/tcp
CMD ["npm", "start"]