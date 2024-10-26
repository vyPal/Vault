import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
	const session = await event.locals.auth();

	let clientlist = await fetch('http://localhost:8080/files/listclients', {
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
		files: fetch(`http://localhost:8080/files/list/${session?.username}/${client}`, {
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

	let vaultFiles = fetch(`http://localhost:8080/files/list/${session?.username}/vault/`, {
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
		vaultFiles: vaultFiles,
		error: null,
		status: 200
	};
};
