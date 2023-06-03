FROM node:alpine
WORKDIR /app
COPY . .
RUN npm install --production
EXPOSE 3000
ENTRYPOINT ["node", "server.js"]