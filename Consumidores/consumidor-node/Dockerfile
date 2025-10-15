FROM node:18-alpine
WORKDIR /app
COPY package.json package.json
RUN npm install
COPY consumer.js .
CMD ["node","consumer.js"]
