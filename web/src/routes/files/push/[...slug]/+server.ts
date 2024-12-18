import { error } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import { SERVER_URL } from '$env/static/private'

export const POST: RequestHandler = async (event) => {
	let user = await event.locals.auth();
	if (!user) {
		return error(401, 'Unauthorized');
	}
	let path = event.params.slug;
	const formData: unknown = Object.fromEntries(await event.request.formData());
	const { file } = formData as { file: File };
	let query = '';
	if (file === undefined) {
		query = '?folder=true';
	}

	let res = await fetch(SERVER_URL + '/files/upload/' + path + query , {
		method: 'POST',
		headers: {
			'Content-Type': 'file',
			Authorization: 'Bearer ' + user.accessToken
		},
		body: file
	});

	return new Response(res.body, res);
};
