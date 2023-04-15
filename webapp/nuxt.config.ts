// https://nuxt.com/docs/api/configuration/nuxt-config

/*import {fileURLToPath} from "url";*/

export default defineNuxtConfig({
    typescript: {
        shim: false
    },
    alias: {
        /*"@": fileURLToPath(new URL("./", import.meta.url)),*/
        assets: "/<rootDir>/assets",
    }
})
