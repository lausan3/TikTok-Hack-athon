FROM node:18-alpine AS base

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install

COPY . .

# Build the app
RUN npm run build

ENV NODE_ENV=production
ENV NEXT_PUBLIC_SERVER_URL=http://hack.tiktok.marouaneh.me
ENV NEXT_PUBLIC_HOST=hack.tiktok.marouaneh.me
# Uncomment the following line in case you want to disable telemetry during runtime.
# ENV NEXT_TELEMETRY_DISABLED 1

EXPOSE 3001

# router pages
CMD ["npm", "start"]