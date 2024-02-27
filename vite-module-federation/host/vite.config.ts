import { defineConfig, loadEnv } from 'vite';
import react from '@vitejs/plugin-react-swc';
import federation from '@originjs/vite-plugin-federation';

interface Imodules {
  remote: string;
}

export default ({ mode }) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };
  const modules: Imodules = {
    remote: '/remote/assets/remoteEntry.js',
  };
  if (mode === 'development') {
    modules.remote = `http://localhost:3002${modules.remote}`;
  }
  return defineConfig({
    plugins: [
      react(),
      federation({
        name: 'host-app',
        remotes: {
          remote: modules.remote,
        },
        shared: ['react'],
      }),
    ],
    build: {
      target: 'esnext',
    },
  });
};
