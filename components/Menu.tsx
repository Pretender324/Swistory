import {
  Card,
  CardContent,
  createStyles,
  makeStyles,
  Theme,
  Typography,
} from '@material-ui/core'
import { Menu } from 'entities/Menu'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: '100%',
    },
    title: {
      paddingLeft: '48px',
    },
    detail: {
      padding: '0 12px 12px',
      float: 'right',
    },
  })
)

export function MenuHeading({ menu }: { menu: Partial<Menu> }) {
  const classes = useStyles()

  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography color="textSecondary">{menu.group?.name}</Typography>
        <Typography variant="h3" className={classes.title}>
          {menu.createdAt}
        </Typography>
        <Typography>
          <span className={classes.detail}>level: {menu.level}</span>
          <span className={classes.detail}>aim: {menu.aim}</span>
        </Typography>
      </CardContent>
    </Card>
  )
}
