import { createStore } from "vuex"
import persist from 'vuex-persistedstate';

export default createStore({
    state: {
        novelInformation: [],
        userInformation: {

        },
        img: '',
        showname: true,

    },
    getters: {
        doneTodos(state) {
            return state.todos.filter(todo => todo.done)
        }
    },

    mutations: {
        increment(state) {
            // εζ΄ηΆζ
            state.count++
        }
    },

    plugins: [
        new persist({
            storage: window.localStorage
        })
    ]
})

