import { createI18n } from 'vue-i18n'
import { defaultLocale } from '@/constant/config';

const localPathPrefix = "../locales/";
const messages = Object.fromEntries(
    Object.entries(import.meta.globEager("../locales/*.json")).map(
        ([key, value]) => {
            const yaml = key.endsWith(".json");
            return [key.slice(localPathPrefix.length, yaml ? -5 : -4), value.default];
        }
    )
);
const locale =  defaultLocale;

export const i18n = createI18n({
    legacy: false,
    locale: locale,
    fallbackLocale: defaultLocale,
    globalInjection: true,
    messages
})

export const t = i18n.global.t;