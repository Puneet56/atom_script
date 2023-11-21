import axios from 'axios';

class CodeService {
	static Instance() {
		return new CodeService();
	}

	async tokenizeCode(code: string) {
		return axios.post(process.env.NEXT_PUBLIC_API_URL + '/api/tokenize', { code });
	}

	async parseCode(code: string) {
		return axios.post(process.env.NEXT_PUBLIC_API_URL + '/api/parse', { code });
	}

	async evaluateCode(code: string) {
		return axios.post(process.env.NEXT_PUBLIC_API_URL + '/api/eval', { code });
	}
}

export default CodeService.Instance();
