import { defineConfig } from "vite";

// The build writes straight into the Go embed directory. That directory is
// gitignored and emptied on every build, so it is never a source of truth.
export default defineConfig({
  base: "./",
  build: {
    outDir: "../internal/viewer/dist",
    emptyOutDir: true,
  },
  server: {
    // Proxy to a locally running arch viewer so dev mode reads real artifacts.
    proxy: {
      "/api": "http://127.0.0.1:7373",
    },
  },
});
