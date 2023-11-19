import { db } from '$lib/server/db';
import { json } from '@sveltejs/kit';

export async function GET() {
    const result = await db.query.usersTable.findFirst({
        with: {
            words: true,
        },
    });

    return json(result);
}