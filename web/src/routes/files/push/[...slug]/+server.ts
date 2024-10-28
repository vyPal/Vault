import { error } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const POST: RequestHandler = async (event) => {
	let user = await event.locals.auth();
	if (!user) {
		return error(401, 'Unauthorized');
	}
	let path = event.params.slug;
	const formData: unknown = Object.fromEntries(await event.request.formData());
	const { file } = formData as { file: File };

	let res = await fetch('http://localhost:8080/files/upload/' + path, {
		method: 'POST',
		headers: {
			'Content-Type': 'file',
			Authorization: 'Bearer ' + user.accessToken
		},
		body: file
	});

	return new Response(res.body, res);
};
