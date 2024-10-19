import { SvelteKitAuth, type DefaultSession } from '@auth/sveltekit';
import Authentik from '@auth/sveltekit/providers/authentik';

declare module '@auth/sveltekit' {
	interface Session {
		accessToken: string;
		username: string;
	}
	interface JWT {
		accessToken: string;
		username: string;
	}
}

export const { handle, signIn, signOut } = SvelteKitAuth({
	providers: [
		Authentik({
			authorization:
				'https://mimlex.vypal.me/application/o/authorize/?scope=openid profile email phone address files:all files:read files:write files:list'
		})
	],
	pages: {
		signIn: '/signin',
		signOut: '/'
	},
	callbacks: {
		async session({ session, token }) {
			if (token.accessToken) {
				session.accessToken = token.accessToken;
			}
			if (token.username) {
				session.username = token.username;
			}
			return session;
		},
		async jwt({ token, account, profile }) {
			if (account) {
				token.accessToken = account.access_token;
			}
			if (profile) {
				token.username = profile.preferred_username;
			}
			return token;
		}
	}
});
