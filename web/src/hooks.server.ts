import { sequence } from '@sveltejs/kit/hooks';

import { prepareStylesSSR } from '@svelteuidev/core';
import { handle as authHandle } from "./auth";

export const handle = sequence(authHandle, prepareStylesSSR);
