import { checkResponse } from '../utils/checkResponse'
import { BASE_URL } from '../utils/constants'

class Api {
  private readonly url: string

  constructor(url: string, slug: string) {
    this.url = url + slug
  }

  async getInfo(data: string): Promise<void> {
    const res = await fetch(`http://localhost/api/status?domain=${data}`, {
      method: 'GET',
    })
    return await checkResponse(res)
  }
}

export const api = new Api(BASE_URL, '/api/status')
