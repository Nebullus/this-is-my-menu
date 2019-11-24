import React from 'react';
import { BrowserRouter, Switch, Route, Redirect } from 'react-router-dom';

import SignInSide from './pages/Home/index';
import SignUp from './pages/SignUp';

export default function Routes() {
    return (
        <BrowserRouter>
            <Switch>
                <Route path="/" exact component={SignInSide} />
                <Route path="/signup" component={SignUp} />
                <Redirect from="*" to="/" />
            </Switch>
        </BrowserRouter>
    );
}