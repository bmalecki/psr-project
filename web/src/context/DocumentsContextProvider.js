import React, { useReducer, createContext, useEffect } from "react";
import { url } from '../environment'

export const DocumentsContext = createContext();

const initialState = [];

function reducer(state, action) {
    switch (action.type) {
        case 'ADD_ITEM':
            return [action.item, ...state];
        case 'REFRESH_FROM_SERVER':
            return [...action.items];
        default:
            throw new Error();
    }
}

export const DocumentsContextProvider = props => {
    const [documents, dispatch] = useReducer(reducer, initialState);

    useEffect(() => {
        async function getDocuments() {
            try {
                const res = await fetch(`${url}/document`)

                const status = await res.status
                if (status !== 200) {
                    console.log("Error");
                }

                const json = await res.json()
                dispatch({ type: "REFRESH_FROM_SERVER", items: json })
            }
            catch (e) {
                console.log("Error getting documents!");
            }
        }

        getDocuments()
    }, []);

    return (
        <DocumentsContext.Provider value={{documents, dispatch}}>
            {props.children}
        </DocumentsContext.Provider>
    );
};