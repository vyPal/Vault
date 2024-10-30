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
		};
	}

	let cl = await clientlist.json();
	let clients = cl.filter((client: string) => client !== 'vault');

	return {
		session: session,
		clientlist: clients,
		status: 200
	};
};
