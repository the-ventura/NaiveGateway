import axios from 'axios'

const url = process.env.NODE_ENV === 'development' ? process.env.REACT_APP_API_URL : window.config.API_URL

const axInstance = axios.create({
  baseURL: url,
  timeout: 1000
})

export default axInstance
