# Imagen base de Node
FROM node:18-alpine

# Setear directorio de trabajo
WORKDIR /app

# Copiar package.json y package-lock.json
COPY package.json package-lock.json ./

# Instalar dependencias
RUN npm install

# Copiar resto del código
COPY . .

# Build de producción
RUN npm run build

# Exponer puerto
EXPOSE 3000

# CMD para correr Next.js en producción
CMD ["npm", "run", "start"]
