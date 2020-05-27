import React, { useReducer, createContext, useEffect } from "react";
import { url } from '../environment'
import { useInterval } from "../intervalHook";

export const DocumentsContext = createContext();

const FAST_REFRESHING = 10000;
const SLOW_REFRESHING = 60000;

const initialState = {
    uploading: false,
    documents: [],
    refreshing: SLOW_REFRESHING
};

function reducer(state, action) {
    switch (action.type) {
        case 'ADD_ITEM':
            return {
                ...state,
                refreshing: FAST_REFRESHING,
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
        case 'INCREASE_REFRESH_RATE':
            return {
                ...state,
                refreshing: FAST_REFRESHING
            }
        case 'DECREASE_REFRESH_RATE':
            return {
                ...state,
                refreshing: SLOW_REFRESHING
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

    async function AsyncGetDocuments() {
        const documents = await getDocuments()
        dispatch({ type: "REFRESH_FROM_SERVER", documents })

        const notAllReady = documents
            .map(i => i["ImageStatus"])
            .some(s => s !== "READY");

        if (notAllReady) {
            dispatch({ type: "INCREASE_REFRESH_RATE" })
        }
        if (state.refreshing === FAST_REFRESHING && notAllReady === false) {
            dispatch({ type: "DECREASE_REFRESH_RATE" })
        }
    }

    useEffect(() => {
        AsyncGetDocuments()
        // eslint-disable-next-line 
    }, []);


    useInterval(() => {
        AsyncGetDocuments()
    }, state.refreshing)

    return (
        <DocumentsContext.Provider value={{ state, dispatch }}>
            {props.children}
        </DocumentsContext.Provider>
    );
};