import { defineStore } from 'pinia'
import { fetchWrapper } from '@/helpers'
import { baseApiUrl } from '../main'

export const useRoomsStore = defineStore({
    id: 'rooms',
    state: () => ({
        rooms: {},
        room: {}
    }),
    actions: {
        async getAll(theaterId) {
            this.rooms = { loading: true };
            try {
                this.thearoomsters = await fetchWrapper.get(`${baseApiUrl}/api/theaters/${theaterId}/rooms`);    
            } catch (error) {
                this.rooms = { error };
            }
        },
    }
});
