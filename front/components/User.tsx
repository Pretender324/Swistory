import {
    Avatar,
    Container,
    createStyles,
    Grid,
    makeStyles,
    Theme,
    Typography,
} from '@material-ui/core'
import React from 'react'

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            textAlign: 'center',
        },
        small: {
            width: theme.spacing(3),
            height: theme.spacing(3),
        },
        large: {
            margin: '0 auto 20px',
            width: theme.spacing(10),
            height: theme.spacing(10),
        },
        title: {
            textAlign: 'left',
        },
        content: {
            marginBottom: '20px',
        },
    })
)

export default function User() {
    const classes = useStyles()
    return (
        <Container maxWidth="lg" className={classes.root}>
            <Avatar className={classes.large}></Avatar>
            <Typography variant="h2">"User name"</Typography>
        </Container>
    )
}
