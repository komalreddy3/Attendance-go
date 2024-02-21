#FROM ubuntu:latest
#LABEL authors="komalreddy"
#
#ENTRYPOINT ["top", "-b"]
# Use a Node.js image
FROM node:latest

# Set the working directory
WORKDIR /frontend

# Copy the package.json and package-lock.json files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy everything to the container
COPY . .

ARG react_app_B_ADDR
ENV react_app_B_ADDR $react_app_B_ADDR

ARG react_app_addr
ENV react_app_addr $react_app_addr
# Expose port
EXPOSE 8000

# Define the command to run the app
CMD ["npm","start"]