import { resolve } from 'path';
import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';
import devtools from 'solid-devtools/vite';

export default defineConfig({
  resolve: {
    alias: {
      '@interfaces': resolve(__dirname, './src/interfaces'),
      '@components': resolve(__dirname, './src/components'),
      '@redux': resolve(__dirname, './src/redux'),
      '@pages': resolve(__dirname, './src/pages'),
      '@styles': resolve(__dirname, './src/styles'),
      '@config': resolve(__dirname, './src/config'),
      '@src': resolve(__dirname, './src'),
      '@assets': resolve(__dirname, './assets'),
      '@hooks': resolve(__dirname, './src/utils/hooks'),
      '@static': resolve(__dirname, './src/static'),
      '@utils': resolve(__dirname, './src/utils'),
    },
  },
  plugins: [
    devtools(),
    solidPlugin(),
  ],
  server: {
    port: 3000,
    host: true,
    strictPort: true,
  },
  build: {
    target: 'esnext',
  },
});
