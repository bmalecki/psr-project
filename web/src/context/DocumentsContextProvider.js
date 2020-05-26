import React, { useReducer, createContext, useEffect } from "react";
import { url } from '../environment'

export const DocumentsContext = createContext();

const initialState = {
    uploading: false,
    documents: []
};

function reducer(state, action) {
    switch (action.type) {
        case 'ADD_ITEM':
            return {
                ...state,
                documents: [action.item, ...state.documents]
            };
        case 'REFRESH_FROM_SERVER':
            return {
                ...state,
                documents: [...action.documents]
            };
        case 'UPLOADING_START':
            return {
                ...state,
                uploading: true
            }
        case 'UPLOADING_END':
            return {
                ...state,
                uploading: false
            }
        default:
            throw new Error();
    }
}

async function getDocuments() {
    try {
        const res = await fetch(`${url}/document`)

        const status = await res.status
        if (status !== 200) {
            console.log("Error");
        }

        return await res.json()
    }
    catch (e) {
        console.log("Error getting documents!");
    }
}

export const DocumentsContextProvider = props => {
    const [state, dispatch] = useReducer(reducer, initialState);

    useEffect(() => {
        async function AsyncFunction() {
            const documents = await getDocuments()
            dispatch({ type: "REFRESH_FROM_SERVER", documents })
        }

        AsyncFunction()
    }, []);

    return (
        <DocumentsContext.Provider value={{ state, dispatch }}>
            {props.children}
        </DocumentsContext.Provider>
    );
};