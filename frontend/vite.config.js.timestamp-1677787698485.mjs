// vite.config.js
import { defineConfig } from "file:///home/parallels/learninggo/AcademicTeamManager/frontend/node_modules/vite/dist/node/index.js";
import vue from "file:///home/parallels/learninggo/AcademicTeamManager/frontend/node_modules/@vitejs/plugin-vue/dist/index.mjs";
var vite_config_default = defineConfig({
  plugins: [vue()],
  server: {
    host: "127.0.0.1",
    port: 8061
  },
  css: {
    preprocessorOptions: {
      less: {
        modifyVars: {
          "primary-color": "#476FFF",
          "link-color": "#476FFF",
          "border-radius-base": "5px"
        },
        javascriptEnabled: true
      }
    }
  }
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcuanMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCIvaG9tZS9wYXJhbGxlbHMvbGVhcm5pbmdnby9BY2FkZW1pY1RlYW1NYW5hZ2VyL2Zyb250ZW5kXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ZpbGVuYW1lID0gXCIvaG9tZS9wYXJhbGxlbHMvbGVhcm5pbmdnby9BY2FkZW1pY1RlYW1NYW5hZ2VyL2Zyb250ZW5kL3ZpdGUuY29uZmlnLmpzXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ltcG9ydF9tZXRhX3VybCA9IFwiZmlsZTovLy9ob21lL3BhcmFsbGVscy9sZWFybmluZ2dvL0FjYWRlbWljVGVhbU1hbmFnZXIvZnJvbnRlbmQvdml0ZS5jb25maWcuanNcIjtpbXBvcnQgeyBkZWZpbmVDb25maWcgfSBmcm9tICd2aXRlJ1xuaW1wb3J0IHZ1ZSBmcm9tICdAdml0ZWpzL3BsdWdpbi12dWUnXG5cbi8vIGh0dHBzOi8vdml0ZWpzLmRldi9jb25maWcvXG5leHBvcnQgZGVmYXVsdCBkZWZpbmVDb25maWcoe1xuICBwbHVnaW5zOiBbdnVlKCldLFxuICBzZXJ2ZXI6IHtcbiAgICBob3N0OiAnMTI3LjAuMC4xJyxcbiAgICBwb3J0OiA4MDYxXG4gIH0sXG4gIGNzczoge1xuICAgIHByZXByb2Nlc3Nvck9wdGlvbnM6IHtcbiAgICAgICAgICBsZXNzOiB7XG4gICAgICAgICAgICAgIG1vZGlmeVZhcnM6IHtcbiAgICAgICAgICAgICAgICAgICdwcmltYXJ5LWNvbG9yJzogJyM0NzZGRkYnLFxuICAgICAgICAgICAgICAgICAgJ2xpbmstY29sb3InOiAnIzQ3NkZGRicsXG4gICAgICAgICAgICAgICAgICAnYm9yZGVyLXJhZGl1cy1iYXNlJzogJzVweCdcbiAgICAgICAgICAgICAgfSxcbiAgICAgICAgICAgICAgamF2YXNjcmlwdEVuYWJsZWQ6IHRydWVcbiAgICAgICAgICB9XG4gICAgICB9XG4gIH0sXG59KVxuIl0sCiAgIm1hcHBpbmdzIjogIjtBQUF1VixTQUFTLG9CQUFvQjtBQUNwWCxPQUFPLFNBQVM7QUFHaEIsSUFBTyxzQkFBUSxhQUFhO0FBQUEsRUFDMUIsU0FBUyxDQUFDLElBQUksQ0FBQztBQUFBLEVBQ2YsUUFBUTtBQUFBLElBQ04sTUFBTTtBQUFBLElBQ04sTUFBTTtBQUFBLEVBQ1I7QUFBQSxFQUNBLEtBQUs7QUFBQSxJQUNILHFCQUFxQjtBQUFBLE1BQ2YsTUFBTTtBQUFBLFFBQ0YsWUFBWTtBQUFBLFVBQ1IsaUJBQWlCO0FBQUEsVUFDakIsY0FBYztBQUFBLFVBQ2Qsc0JBQXNCO0FBQUEsUUFDMUI7QUFBQSxRQUNBLG1CQUFtQjtBQUFBLE1BQ3ZCO0FBQUEsSUFDSjtBQUFBLEVBQ0o7QUFDRixDQUFDOyIsCiAgIm5hbWVzIjogW10KfQo=
