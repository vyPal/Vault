import { SvelteKitAuth } from "@auth/sveltekit"
import Authentik from "@auth/sveltekit/providers/authentik"
 
export const { handle } = SvelteKitAuth({
  providers: [Authentik({
    authorization: "https://mimlex.vypal.me/application/o/authorize/?scope=\"openid profile email phone address files:all files:read files:write files:list\""
  })],
})
