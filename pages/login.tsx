/* eslint-disable indent */
import {
  Button,
  createStyles,
  Grid,
  makeStyles,
  TextField,
  Theme,
  Typography,
} from '@material-ui/core'
import React, { useState } from 'react'

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

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [email, setEmail] = useState<string>()
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [password, setPassword] = useState<string>()

  const handleChange = (event) => {
    console.log(event.target.name)
    switch (event.target.name) {
      case 'email':
        setEmail(event.target.value)
        break
      case 'password':
        setPassword(event.target.value)
        break
      default:
        console.error('key not found')
    }
  }

  const handleSubmit = async (event) => {
    event.preventDefault()
    const res = await fetch('localhost:5000/v1/user/login', {
      body: JSON.stringify({
        email: email,
        password: password,
      }),
      headers: {
        'Content-Type': 'application/json',
      },
      method: 'POST',
    })
    const result = await res.json()
    console.log(result)
  }

  return (
    <div className={classes.root}>
      <Typography variant="h1" className={classes.title}>
        Swistory
      </Typography>
      <Grid container justify="center">
        <Grid item xs={8} md={4}>
          <form onSubmit={handleSubmit}>
            <TextField
              variant="filled"
              fullWidth
              placeholder="E-mail"
              margin="normal"
              onChange={handleChange}
            ></TextField>
            <TextField
              variant="filled"
              fullWidth
              placeholder="Password"
              margin="normal"
              onChange={handleChange}
            ></TextField>
            <Button
              variant="contained"
              fullWidth
              className={classes.button}
              type="submit"
            >
              Login
            </Button>
          </form>
        </Grid>
      </Grid>
    </div>
  )
}
