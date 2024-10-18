import type { PageServerLoad } from "./$types"

export const load: PageServerLoad = async (event) => {
  const session = await event.locals.auth()

  let clientlist = await fetch('http://localhost:8080/files/listclients', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + session?.accessToken
    }
  })

  if (clientlist.status !== 200) {
    return {
      status: clientlist.status,
      error: await clientlist.text()
    }
  } else {
    let nl = []
    for (let client of await clientlist.json()) {
      let clientFiles = await fetch('http://localhost:8080/files/list/' + session?.username + '/' + client, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ' + session?.accessToken
        }
      })
      if (clientFiles.status !== 200) {
        return {
          status: clientFiles.status,
          error: await clientFiles.text()
        }
      } else {
        nl.push({
          name: client,
          files: await clientFiles.json()
        })
      }
    }
    return {
      session: session,
      clientlist: nl
    }
  }
}
