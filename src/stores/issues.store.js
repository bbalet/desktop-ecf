import { defineStore } from 'pinia'
import { fetchWrapper } from '@/helpers'
import { baseApiUrl } from '../main'

export const useIssuesStore = defineStore({
    id: 'issues',
    state: () => ({
        issues: {},
        issue: {}
    }),
    actions: {
        async getAll(roomId) {
            this.issues = { loading: true };
            try {
                this.issues = await fetchWrapper.get(`${baseApiUrl}/api/rooms/${roomId}/issues/`);    
            } catch (error) {
                this.issues = { error };
            }
        },
        async getById(id) {
            this.issue = { loading: true };
            try {
                this.issue = await fetchWrapper.get(`${baseApiUrl}/api/issues/${id}`);
            } catch (error) {
                this.issue = { error };
            }
        },
        async update(id, params) {
            await fetchWrapper.patch(`${baseApiUrl}/api/issues/${id}`, params);

            
        },
        async create(roomId, params) {
            await fetchWrapper.post(`${baseApiUrl}/api/rooms/${roomId}/issues`, params);

            
        },
    }
});
