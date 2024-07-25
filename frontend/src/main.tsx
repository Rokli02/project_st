import React from 'react'
import {createRoot} from 'react-dom/client'
import './style.css'
import Layout from './Layout'
import { HashRouter } from 'react-router-dom'
import { AuthProvider } from './context'

const root = createRoot(document.getElementById('root')!)

root.render(
    <React.StrictMode>
        <HashRouter basename='/'>
            <AuthProvider>
                <Layout/>
            </AuthProvider>
        </HashRouter>
    </React.StrictMode>
)
