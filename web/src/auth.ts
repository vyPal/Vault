import { SvelteKitAuth, type DefaultSession } from "@auth/sveltekit"
import Authentik from "@auth/sveltekit/providers/authentik"

declare module "@auth/sveltekit" {
  interface Session {
    accessToken: string
  }
  interface JWT {
    accessToken: string
  }
}
 
export const { handle, signIn, signOut } = SvelteKitAuth({
  providers: [Authentik({
    authorization: "https://mimlex.vypal.me/application/o/authorize/?scope=\"openid profile email phone address files:all files:read files:write files:list\""
  })],
  pages: {
    signIn: "/signin",
    signOut: "/",
  },
  callbacks: {
    async session({ session, token }) {
      console.log(token)
      if (token.accessToken) {
        session.accessToken = token.accessToken
      }
      console.log(session)
      return session
    },
    async jwt({ token, account }) {
      console.log(account)
      if (account) {
        token.accessToken = account.access_token
      }
      console.log(token)
      console.log("and then")
      return token
    }
  }
})
