# pull official base image
FROM node:12.18.3-alpine

# set working directory
WORKDIR /src

# add `/app/node_modules/.bin` to $PATH
ENV PATH /src/node_modules/.bin:$PATH

# install app dependencies
COPY package*.json ./
RUN npm install --silent

# add app
COPY . ./

# start app
CMD ["npm", "start"]