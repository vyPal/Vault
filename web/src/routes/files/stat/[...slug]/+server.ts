import { error } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { SERVER_URL } from '$env/static/private'

export const GET: RequestHandler = async (event) => {
	let user = await event.locals.auth();
	if (!user) {
		return error(401, 'Unauthorized');
	}
	let path = event.params.slug;
	let query = '';
	if (event.url.searchParams.get('folder') === 'true') {
		query = '?folder=true';
	}
	let res = await fetch(SERVER_URL + '/files/metadata/' + path + query, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
			Authorization: 'Bearer ' + user.accessToken
		}
	});

	return new Response(res.body, res);
};
