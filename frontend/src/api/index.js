import axios from 'axios';

var HTTP = axios.create({
  timeout: 10000,
  responseType:'json',
  headers:{
   
  },
  
})

export default HTTP