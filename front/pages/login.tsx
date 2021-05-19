import classes from '*.module.css'
import {
    Button,
    createStyles,
    Grid,
    makeStyles,
    TextField,
    Theme,
    Typography,
} from '@material-ui/core'
import React from 'react'

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            textAlign: 'center',
            minHeight: '100vh',
        },
        title: {
            paddingTop: '30vh',
        },
        button: {
            marginTop: '20px',
        },
    })
)

export default function Login() {
    const classes = useStyles()
    return (
        <div className={classes.root}>
            <Typography variant="h1" className={classes.title}>
                Swistory
            </Typography>
            <Grid container justify="center">
                <Grid item xs={8} md={4}>
                    <form action="">
                        <TextField
                            variant="filled"
                            fullWidth
                            placeholder="Username"
                            margin="normal"
                        ></TextField>
                        <TextField
                            variant="filled"
                            fullWidth
                            placeholder="Password"
                            margin="normal"
                        ></TextField>
                        <Button
                            variant="contained"
                            fullWidth
                            className={classes.button}
                        >
                            Login
                        </Button>
                    </form>
                </Grid>
            </Grid>
        </div>
    )
}
