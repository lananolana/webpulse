import { checkResponse } from '../utils/checkResponse'
import { BASE_URL } from '../utils/constants'

class Api {
  private readonly url: string

  constructor(url: string, slug: string) {
    this.url = url + slug
  }

  async getInfo(data: string): Promise<void> {
    //const req = {'path': data}
    const res = await fetch(`http://localhost/api/status?domain=${data}`, {
      //method: 'GET',
      mode: 'no-cors',
      //sec-fetch-site: 'cross-site',
      //referrerPolicy: "strict-origin-when-cross-origin",
      /*headers: {
        'Content-Type': 'application/json',
      },
      */credentials: 'include',
      /*headers: {
        'Content-Type': 'application/json',
      },*/
      //credentials: 'include',
    })
    return await /*checkResponse(*/res.json()/*)*/
  }

  /*
    async getInfo(data: string): Promise<void> {
    //const req = {'path': data}
    const res = await fetch(`${this.url}`, {
      method: '',
      mode: 'no-cors',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(/*req*//*data),
      credentials: 'include',
    })
    return await checkResponse(res)
  }
  */
}

export const api = new Api(BASE_URL, '/api/status')
