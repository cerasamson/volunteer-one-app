import React from 'react';
import PropTypes from 'prop-types';
import { StyleSheet, Image } from 'react-native';
import { Block, Text, theme } from 'galio-framework';


class PostWithoutImage extends React.Component {
  render() {
    const {item, horizontal, style} = this.props;
    
    const cardContainer = [styles.card, styles.shadow, style];
    const profileImageContainer = [styles.profileImageContainer,
      styles.shadow
    ];

    return (
      <Block row={horizontal} card flex style={cardContainer}>
        {/* Profile Image */}
        <Block style={[ styles.cardContainer, { flexDirection: 'column', } ]}>  
        {/* Displays info. about the post author. */}
          <Block style={[profileImageContainer , {flexDirection: 'row', } ]}>
            <Image source={{uri: item.profileImage}} style={styles.fullImage} />
            <Block style={styles.cardTitleContainer}>            
              <Text size={18} bold style={styles.cardTitle}>{item.author}</Text>
              <Text size={14} style={styles.cardTitle}>{item.date}</Text>
            </Block>
          </Block>  
        {/* Displays user description. */}
            <Block style={styles.postDescriptionContainer}>
              <Text size={14} style={styles.cardDescription}>{item.description}</Text>
            </Block>
          </Block>
      </Block>
    );
  }
}

PostWithoutImage.propTypes = {
  item: PropTypes.object,
  horizontal: PropTypes.bool,
  full: PropTypes.bool,
  ctaColor: PropTypes.string,
  imageStyle: PropTypes.any,
}

const styles = StyleSheet.create({
  card: {
    backgroundColor: theme.COLORS.WHITE,
    marginVertical: theme.SIZES.BASE,
    borderWidth: 0,
    minHeight: 150,
    marginBottom: 16
  },
  cardTitleContainer: {
    paddingLeft: 10,
    paddingTop: 20
  },
  cardDescription: {
    padding: theme.SIZES.BASE / 2
  },
  profileImageContainer: {
    paddingLeft: 10,
    paddingTop: 10
  },
  fullImage: {
    width: 70,
    height: 70,
    borderRadius: 70/2
  },
  postDescriptionContainer: {
    margin: theme.SIZES.BASE,
  },
  shadow: {
    shadowColor: theme.COLORS.BLACK,
    shadowOffset: { width: 0, height: 2 },
    shadowRadius: 4,
    shadowOpacity: 0.1,
    elevation: 2,
  },
});

export default PostWithoutImage;