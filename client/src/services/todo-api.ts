import { Configuration } from '../api/api-client';
import { TodosApi } from '../api/api-client/api';

const config = new Configuration({
  basePath: import.meta.env.VITE_API_BASE_URL,
});

export default new TodosApi(config);
