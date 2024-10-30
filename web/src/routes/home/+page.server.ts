import type { PageServerLoad } from './$types';
import { SERVER_URL } from '$env/static/private'

export const load: PageServerLoad = async (event) => {
	const session = await event.locals.auth();

	let clientlist = await fetch(SERVER_URL + '/files/listclients', {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
			Authorization: 'Bearer ' + session?.accessToken
		}
	});

	if (clientlist.status !== 200) {
		return {
			status: clientlist.status,
			error: await clientlist.text(),
			session: session,
			clientlist: [],
			vaultFiles: []
		};
	}

	let clients = await clientlist.json();
	clients = clients.filter((client: string) => client !== 'vault');
	let nl = clients.map((client: string) => ({
		name: client,
		files: fetch(SERVER_URL + `/files/list/${session?.username}/${client}`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				Authorization: 'Bearer ' + session?.accessToken
			}
		})
			.then(async (response) => {
				if (response.status !== 200) {
					throw new Error(await response.text());
				}
				return response.json();
			})
			.catch((error) => ({
				error: error.message
			}))
	}));

	let vaultFiles = fetch(SERVER_URL + `/files/list/${session?.username}/vault/`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
			Authorization: 'Bearer ' + session?.accessToken
		}
	}).then(async (response) => {
		if (response.status !== 200) {
			throw new Error(await response.text());
		}
		return response.json();
	}).catch((error) => ({
		error: error.message
	}));

	return {
		session: session,
		clientlist: nl ?? [],
		vaultFiles: await vaultFiles,
		error: null,
		status: 200
	};
};
