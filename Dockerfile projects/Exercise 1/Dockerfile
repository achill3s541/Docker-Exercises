FROM node:19.0.0
WORKDIR /usr/app
COPY package.json package-lock.json ./
RUN npm install
#Adding custom user&group - appjs
RUN groupadd -g 1001 appjs && useradd -u 1001 -g appjs appjs && mkdir /home/appjs
#Adding priviliges for custom users
RUN chown root:appjs -R ./ && chmod 750 -R ./* && chown appjs:appjs -R /home/appjs && chmod 750 -R /home/appjs
USER appjs:appjs
COPY src ./src
EXPOSE 3030/tcp
CMD ["npm", "start"]