// STYLES

// EXTERNALS
import dot from 'dot-object';

// LOCALS
import enUs from './en_US.json';
import frFR from './fr_FR.json';
import { INTL } from '../const/intl/intl';
import { useConfigStore } from '../store/store';

type Copy = Record<string, string | object>;

class I18n {
  private lang: string;
  private copy: Copy;

  constructor() {
    this.lang = this._getLang();
    this.copy = this._getCopy();
  }

  public t(key: string): string {
    let copy = this.copy;
    // we are using pick in order to make life easier when the day for plurals and copy with params arrives
    const result = dot.pick(key, copy);
    return result ?? key;
  }

  private _getLang(): string {
    let urlParams = new URLSearchParams(window.location.search);
    let fromUrl = urlParams.get('l');

    if (fromUrl) {
      useConfigStore.setState({ lang: fromUrl });
    }

    return fromUrl || useConfigStore.getState().lang || INTL.EN_us;
  }

  private _getCopy(): Copy {
    let lang = this.lang;

    if (lang === INTL.EN_us) {
      return enUs;
    } else if (lang === INTL.FR_fr) {
      return frFR;
    } else {
      console.warn(
        `I18n::_getCopy:: We may not support yet ${lang}. Loading 'en_US' then... `,
      );
      return enUs;
    }
  }
}

const Intl = new I18n();
Object.freeze(Intl);

export default Intl;
