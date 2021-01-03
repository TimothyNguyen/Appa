import React from 'react';

import { Container } from 'react-bootstrap';

const Footer = () => (
    <Container>
        <div 
            style={{
                position: 'absolute',
                bottom: '0',
            }}>
            <div>© 2021 Appa</div>
        </div>
    </Container>
);

export default Footer;